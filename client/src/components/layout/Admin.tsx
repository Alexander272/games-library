import { Suspense } from "react"
import { Outlet } from "react-router-dom"
import { Navbar } from "../Navbar/Navbar"
import "./admin.scss"

export default function Admin() {
    return (
        <div className='admin'>
            <Navbar />
            <div className='main'>
                <Suspense fallback={"Loading..."}>
                    <Outlet />
                </Suspense>
            </div>
        </div>
    )
}
