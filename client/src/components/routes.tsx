import { lazy, Suspense } from "react"
import { Routes, Route } from "react-router-dom"

const Auth = lazy(() => import("../pages/Auth/Auth"))
const NotFound = lazy(() => import("../pages/NotFound/NotFound"))

export const MyRoutes = () => {
    // const { isAuth, role } = useSelector((state: RootState) => state.user)

    let routes = (
        <>
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
