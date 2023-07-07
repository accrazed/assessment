import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'

export default function Login({ setCookie }) {
	const [username, setUsername] = useState('')
	const [password, setPassword] = useState('')

	let navigate = useNavigate()

	const handleLogin = () => {
		fetch(`http://localhost:8080/session`, {
			method: 'POST',
			body: JSON.stringify({
				username: username,
				password: password,
			}),
		}).then((resp) => {
			if (resp.ok) {
				resp.json().then((json) => {
					console.log(json)
					setCookie('UserID', json.UserID)
					navigate('/home')
				})
			}
		})
	}

	return (
		<div className="Login">
			<input
				placeholder="Username"
				onChange={(e) => {
					setUsername(e.target.value)
				}}
			/>
			<input
				placeholder="Password"
				onChange={(e) => {
					setPassword(e.target.value)
				}}
			/>
			<button onClick={handleLogin}>Login</button>
		</div>
	)
}
