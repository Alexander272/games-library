import React, { FC, ReactElement } from "react"
import classes from "./input.module.scss"

interface Props {
    id?: string
    name?: string
    value?: string
    onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
    rules?: Rules[]
    attr?: React.InputHTMLAttributes<HTMLInputElement>
}

type Rules = BaseRule

declare type RuleType =
    | "string"
    | "number"
    | "boolean"
    | "method"
    | "regexp"
    | "integer"
    | "float"
    | "object"
    | "enum"
    | "date"
    | "url"
    | "hex"
    | "email"

interface BaseRule {
    len?: number
    max?: number
    message?: string | ReactElement
    min?: number
    pattern?: RegExp
    required?: boolean
    type?: RuleType
}

export const Input: FC<Props> = ({ id, name, value, onChange, rules, attr }) => {
    return (
        <input
            id={id}
            name={name}
            value={value}
            className={classes.input}
            {...attr}
            onChange={onChange}
        />
    )
}
