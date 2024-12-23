import { useState } from 'react'
import Pages from './pages/Pages.jsx'
import { GetAll } from './http.js'

const App = () => {
    const [page, setPage] = useState('workouts')

    const save = () => {
        GetAll('save').catch((err) => {
            console.log("save error: ", err)
        })
    }

    return (
        <div className='app-grid'>
            <nav className='nav'>
                <div>
                    <button type='button' onClick={save}>Save</button>
                </div>
                <ul>
                    <li>
                        <button type='button' onClick={() => setPage('exercises')}>Exercises</button>
                    </li>
                    <li>
                        <button type='button' onClick={() => setPage('workouts')}>Workouts</button>
                    </li>
                </ul>
            </nav>
            <main className='content'>
                <Pages page={page} />
            </main>
        </div>
    )
}

export default App
