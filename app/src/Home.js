import React, { useState, useEffect } from 'react'

export default function Home({ userID }) {
	const [users, setUsers] = useState([])

	useEffect(() => {
		fetch(`http://localhost:8080/users/${userID}`, {
			method: 'GET',
		}).then((resp) => {
			if (resp.ok) {
				resp.json().then((json) => {
					setUsers([json])
					console.log(json)
				})
			}
		})
	}, [userID])

	// TODO: Add admin functionality to list users

	return (
		<div className="Home">
			Welcome!
			{users.map((user) => (
				<div key={user.userID}>
					<div>
						Name: {user.firstName} {user.lastName}
					</div>
					<div>Username: {user.username} </div>
					<div>Email: {user.email}</div>
					<div>ID: {user.userID}</div>
				</div>
			))}
		</div>
	)
}
