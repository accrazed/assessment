import './App.css'
import { useCookies } from 'react-cookie'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Login from './Login'
import Home from './Home'

function App() {
	const [cookie, setCookie] = useCookies(['UserID'])

	return (
		<div className="App">
			<Router>
				<Routes>
					<Route exact path="/" element={<Login setCookie={setCookie} />} />
					<Route exact path="/" element={<Home userID={cookie.UserID} />} />
				</Routes>
			</Router>
		</div>
	)
}

export default App
