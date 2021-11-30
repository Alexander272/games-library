package storage

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/chai2010/webp"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type FileStore struct {
	storage    *storage.BucketHandle
	bucketName string
}

func NewFileStorage(bucketName, pathToCredentials string) (*FileStore, error) {
	config := firebase.Config{
		StorageBucket: bucketName + ".appspot.com",
	}
	opt := option.WithCredentialsFile(pathToCredentials)
	app, err := firebase.NewApp(context.Background(), &config, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to create new app. error : %w", err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get storage. error: %w", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket. error: %w", err)
	}

	return &FileStore{
		storage:    bucket,
		bucketName: bucketName,
	}, nil
}

func (fs *FileStore) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, path, name string) (File, error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return File{}, fmt.Errorf("failed to read file. error: %w", err)
	}

	var newFile []byte
	var filename string
	contentType := header.Header.Get("Content-Type")
	if contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/jpg" {
		newFile, err = fs.imageCompressing(fileBytes, 85, contentType)
		if err != nil {
			return File{}, fmt.Errorf("falied to compressing file. error: %w", err)
		}

		if name != "" {
			filename = name + ".webp"
		} else {
			filename = fmt.Sprintf("%s_%d.webp", strings.Split(header.Filename, ".")[0], time.Now().Unix())
		}
	} else {
		newFile = fileBytes
		nameParts := strings.Split(header.Filename, ".")
		if name != "" {
			filename = fmt.Sprintf("%s.%s", name, nameParts[1])
		} else {
			filename = fmt.Sprintf("%s_%d.%s", nameParts[0], time.Now().Unix(), nameParts[1])
		}
	}

	uuid := uuid.New()
	wc := fs.storage.Object(filepath.Join(path, filename)).NewWriter(ctx)
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": uuid.String()}
	wc.ObjectAttrs.MediaLink = fmt.Sprintf("https://storage.cloud.google.com/%s.appspot.com/%s", fs.bucketName, filepath.Join(path, filename))

	_, err = io.Copy(wc, bytes.NewReader(newFile))
	if err != nil {
		return File{}, fmt.Errorf("failed to send file. error: %w", err)
	}

	if err = wc.Close(); err != nil {
		return File{}, fmt.Errorf("failed to close connection. error: %w", err)
	}
	if err := fs.storage.Object(filepath.Join(path, filename)).ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return File{}, err
	}

	return File{Name: filename, Url: wc.MediaLink}, nil
}

func (fs *FileStore) Remove(ctx context.Context, path, filename string) error {
	if filename != "" {
		if err := fs.storage.Object(filepath.Join(path, filename)).Delete(ctx); err != nil {
			return fmt.Errorf("failed to delete file. error: %w", err)
		}
		return nil
	}

	objects := fs.storage.Objects(ctx, &storage.Query{Prefix: path})
	obj, err := objects.Next()
	if err != nil && err.Error() != "no more items in iterator" {
		return err
	}
	for obj != nil {
		err = fs.storage.Object(obj.Name).Delete(ctx)
		if err != nil {
			return err
		}
		obj, err = objects.Next()
		if err != nil && err.Error() != "no more items in iterator" {
			return err
		}
	}
	return nil
}

func (fs *FileStore) imageCompressing(buffer []byte, quality float32, contentType string) ([]byte, error) {
	var img image.Image
	var err error
	switch contentType {
	case "image/png":
		img, err = png.Decode(bytes.NewReader(buffer))
		if err != nil {
			return nil, err
		}
	case "image/jpeg", "image/jpg":
		img, err = jpeg.Decode(bytes.NewReader(buffer))
		if err != nil {
			return nil, err
		}
	case "image/webp":
		return buffer, nil
	}

	var out bytes.Buffer
	if err = webp.Encode(&out, img, &webp.Options{Lossless: true, Exact: true, Quality: quality}); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
