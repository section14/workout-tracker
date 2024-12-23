import { useCallback, useEffect, useState } from 'react'
import { Create, Delete, GetAll, Update } from '../http.js'

const ExerciseSelect = ({selected, onChange}) => {
    const [exercises, setExercises] = useState([])
    const [value, setValue] = useState(selected ? selected : 1)
    
    useEffect(() => {
        onChange(value)
    }, [value])

    useEffect(() => {
        GetAll('exercises').then((res) => {
            setExercises(res)
        }).catch((err) => {
            console.log("error getting exercises: ", err)
        })
    }, [])

    return (
        <select value={value} onChange={(e) => setValue(parseInt(e.target.value))}>
            {exercises.map((e) => (
                <option key={e.id} value={e.id}>{e.name}</option>
            ))}
        </select>
    )
}

export default ExerciseSelect
