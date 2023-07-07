import './App.css'
import { useCookies } from 'react-cookie'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Login from './Login'

function App() {
	const [cookie, setCookie] = useCookies(['UserID'])

	return (
		<div className="App">
			<Router>
				<Routes>
					<Route exact path="/" element={<Login setCookie={setCookie} />} />
				</Routes>
			</Router>
		</div>
	)
}

export default App
