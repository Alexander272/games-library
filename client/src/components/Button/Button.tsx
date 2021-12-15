import React, { FC } from "react"
import classes from "./button.module.scss"

export interface ButtunProps {
    children?: React.ReactNode
    onClick: () => void
    size?: "small" | "middle" | "large"
    rounded?: "none" | "low" | "medium" | "high" | "round" | "circle"
    attr?: React.ButtonHTMLAttributes<any>
    // disabled: boolean
}

export const Button: FC<ButtunProps> = ({ children, onClick, size, rounded, attr }) => {
    return (
        <button
            onClick={onClick}
            className={[
                classes.button,
                classes[size || "middle"],
                classes[rounded || "medium"],
            ].join(" ")}
            {...attr}
        >
            {children}
        </button>
    )
}
