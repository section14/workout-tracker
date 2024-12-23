import { useCallback, useEffect, useState } from 'react'
import { Create, Delete, GetAll, Update } from '../http.js'

const CreateExercise = ({ }) => {
    const [nameVal, setNameVal] = useState('')

    const handleCreate = useCallback(() => {
        console.log("create: ", nameVal)
        const req = { name: nameVal }

        Create('exercise', req).catch((err) => {
            console.log("err: ", err)
        })
    }, [nameVal])

    return (
        <form id='create-exercise' style={{ marginBottom: '1.0rem' }} onSubmit={(e) => e.preventDefault()}>
            <input type='text' name='name' onChange={(e) => setNameVal(e.target.value)} value={nameVal} />
            <button type='submit' onClick={handleCreate}>Submit</button>
        </form>
    )
}

const Exercise = ({ id, name }) => {
    const [state, setState] = useState(name)

    const remove = (id) => {
        Delete('exercise', id).catch((err) => {
            console.log("can't delete", err)
        })
    }

    const save = useCallback((id) => {
        const req = {id, name: state}

        Update('exercise', req, id).catch((err) => {
            console.log("can't update", err)
        })
    }, [state])

    return (
        <div>
            <div style={{ display: 'inline-flex' }}>
                <input type='text' onChange={(e) => setState(e.target.value)} value={state} />
                <button type='button' onClick={() => save(id)}>Save</button>
                <button type='button' onClick={() => remove(id)}>Delete</button>
            </div>
        </div>
    )
}

const Exercises = ({ }) => {
    const [exercises, setExercises] = useState([])

    useEffect(() => {
        GetAll('exercises').then((res) => {
            setExercises(res)
        }).catch((err) => {
            console.log("error getting exercises: ", err)
        })
    }, [])

    return (
        <>
            <CreateExercise />
            <div style={{ paddingBottom: '1.0rem' }}>Exercises</div>
            {exercises.map((e) => (
                <Exercise key={e.id} id={e.id} name={e.name} />
            ))}
        </>
    )
}

export default Exercises
