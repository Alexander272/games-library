import { Link } from "react-router-dom"
import classes from "./notFound.module.scss"

export default function NotFound() {
    return (
        <div id='notfound'>
            <div className={classes.notfound}>
                <div className='notfound-404'>
                    <h1>Oops!</h1>
                </div>
                <h2>404 - Page not found</h2>
                <p>
                    The page you are looking for might have been removed had its name changed or is
                    temporarily unavailable.
                </p>
                <Link to='/'>Go To Homepage</Link>
            </div>
        </div>
    )
}
