import { useState } from 'react'
import Pages from './pages/Pages'
import './App.css'

const App = () => {
    const [page, setPage] = useState('workouts')

    return (
        <div className='app-grid'>
            <nav className='nav'>
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
