import { useState } from "react"
import { Form } from "../../components/Form/Form"
import { Input } from "../../components/Input/Input"
import classes from "./auth.module.scss"

export default function Auth() {
    const [tab, setTab] = useState("singIn")

    const changeTab = (tabName: string) => () => setTab(tabName)

    return (
        <div className={classes.auth}>
            <div
                className={[
                    classes.container,
                    tab === "singIn" ? classes.signIn : classes.signUp,
                ].join(" ")}
            >
                <div className={classes.form}>
                    <ul className={classes.nav}>
                        <li className={`${classes.item} ${tab === "singIn" && classes.active}`}>
                            <span onClick={changeTab("singIn")}>Sign In</span>
                        </li>
                        <li className={`${classes.item} ${tab !== "singIn" && classes.active}`}>
                            <span onClick={changeTab("singUp")}>Sign Up</span>
                        </li>
                    </ul>
                    {/* <transition name="slide-fade" mode="out-in">
                    <form class="tab" v-if="isSignIn" @submit.prevent="SignIn(signIn)">
                        <input-field
                            id="email"
                            name="email"
                            type="email"
                            labelText="Email"
                            :errorText="signIn.errorEmail"
                            v-model="signIn.email"
                        />
                        <input-field
                            id="password"
                            name="password"
                            type="password"
                            labelText="Пароль"
                            :errorText="signIn.errorPassword"
                            v-model="signIn.password"
                        />
                        <button class="submit" @click="checkFormSignIn">Sign in</button>
                    </form>
                    <form class="tab" v-else @submit.prevent="SignUp(signUp)">
                        <input-field
                            id="name"
                            name="name"
                            labelText="Логин"
                            :errorText="signUp.errorLogin"
                            v-model="signUp.login"
                        />
                        <input-field
                            id="email"
                            name="email"
                            type="email"
                            labelText="Email"
                            :errorText="signUp.errorEmail"
                            v-model="signUp.email"
                        />
                        <input-field
                            id="password"
                            name="password"
                            type="password"
                            labelText="Пароль"
                            :errorText="signUp.errorPassword"
                            v-model="signUp.password"
                        />
                        <button class="submit" @click="checkformSignUp">Sign up</button>
                    </form>
                </transition> */}
                </div>
                {/* <a href='#' class='forgot'>
                    Forgot Password?
                </a>  */}
            </div>
        </div>
    )
}
