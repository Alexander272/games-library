import { lazy, Suspense } from "react"
import { Routes, Route } from "react-router-dom"

const Admin = lazy(() => import("./layout/Admin"))
const Dashboard = lazy(() => import("../pages/Dashboard/Dashboard"))
const Games = lazy(() => import("../pages/Games/Games"))
const Users = lazy(() => import("../pages/Users/Users"))
const Auth = lazy(() => import("../pages/Auth/Auth"))
const NotFound = lazy(() => import("../pages/NotFound/NotFound"))

export const MyRoutes = () => {
    // const { isAuth, role } = useSelector((state: RootState) => state.user)

    let routes = (
        <>
            <Route path='/admin/' element={<Admin />}>
                <Route index element={<Dashboard />} />
                <Route path='games/' element={<Games />} />
                <Route path='users/' element={<Users />} />
            </Route>
            <Route path='/auth/' element={<Auth />} />
            <Route path='/*' element={<NotFound />} />
        </>
    )

    return (
        <Suspense fallback={"Loading..."}>
            <Routes>{routes}</Routes>
        </Suspense>
    )
}
