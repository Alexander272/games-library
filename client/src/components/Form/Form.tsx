import {
    Children,
    cloneElement,
    PropsWithChildren,
    ReactElement,
    useMemo,
    useEffect,
    FormEvent,
    FC,
    useState,
} from "react"
import { Button } from "../Button/Button"
import classes from "./form.module.scss"

interface Props {
    children: ReactElement[] | ReactElement
}

interface ItemProps {
    label?: string
    id?: string
    name?: string
    initValue?: any
    children: ReactElement[] | ReactElement
}

interface State {
    value: string
    error: null | string
}

const Form = ({ children }: PropsWithChildren<Props>): React.ReactElement => {
    const [state, setState] = useState<State[]>([])

    const providerState = useMemo(
        () => ({
            ...state,
        }),
        [state]
    )

    useEffect(() => {
        console.log("render")
    }, [])

    const submitHandler = (event: FormEvent) => {
        event.preventDefault()
    }

    return (
        <form className={classes.form} onSubmit={submitHandler}>
            {Children.map(children, (child: React.ReactElement) => {
                console.log(child)
                console.log(child.key)

                return cloneElement(child, {
                    ...providerState,
                })
            })}
        </form>
    )
}

const Item: FC<ItemProps> = ({
    label,
    id,
    name,
    initValue,
    children,
}: PropsWithChildren<ItemProps>) => {
    const providerState = useMemo(() => ({}), [])
    return (
        <div className={classes.item}>
            {label && (
                <label htmlFor={id} title={label} className={classes.label}>
                    {label}
                </label>
            )}
            {Children.map(children, (child: React.ReactElement) => {
                console.log(child)

                return cloneElement(child, {
                    ...providerState,
                })
            })}
        </div>
    )
}

Form.Button = Button
Form.Item = Item

export { Form }
