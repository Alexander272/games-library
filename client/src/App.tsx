import { BrowserRouter } from "react-router-dom"
import { MyRoutes } from "./components/routes"

function App() {
    return (
        <BrowserRouter basename={process.env.BASE_URL}>
            <div className='wrapper'>
                <MyRoutes />
            </div>
        </BrowserRouter>
    )
}

export default App
