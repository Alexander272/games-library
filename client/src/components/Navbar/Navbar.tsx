import { FC } from "react"
import { NavLink } from "react-router-dom"
import classes from "./navbar.module.scss"

export const Navbar: FC = () => {
    return (
        <div className={classes.navbar}>
            <div className={classes.user}>
                <h4>
                    <span className={classes.icon}>icon</span>
                    Username
                </h4>
            </div>
            <nav className={classes.nav}>
                <ul className={classes.list}>
                    <li className={classes.item}>
                        <NavLink
                            to='/admin/'
                            key='Dashboard'
                            className={({ isActive }) =>
                                `${classes.link} ${isActive && classes.active}`
                            }
                        >
                            <span className={classes.icon}>icon</span> Dashboard
                        </NavLink>
                    </li>
                    <li className={classes.item}>
                        <NavLink
                            to='/admin/games/'
                            key='Dashboard'
                            className={({ isActive }) =>
                                `${classes.link} ${isActive && classes.active}`
                            }
                        >
                            <span className={classes.icon}>icon</span> Dashboard
                        </NavLink>
                    </li>
                    <li className={classes.item}>
                        <NavLink
                            to='/admin/users/'
                            key='Dashboard'
                            className={({ isActive }) =>
                                `${classes.link} ${isActive && classes.active}`
                            }
                        >
                            <span className={classes.icon}>icon</span> Dashboard
                        </NavLink>
                    </li>
                </ul>
            </nav>
        </div>
    )
}
