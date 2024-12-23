import { useCallback, useEffect, useReducer, useState } from 'react'
import { Create, Delete, Get, GetAll, Update } from '../http.js'
import ExerciseSelect from '../widgets/ExerciseSelect.jsx'

const Action = ({ action, onChange }) => {
    const actionReducer = (state, action) => {
        switch (action.type) {
            case 'set-all':
                return {
                    id: action.payload.id,
                    exerciseId: action.payload.exerciseId,
                    sets: action.payload.sets,
                    reps: action.payload.reps,
                    available: true,
                }
            case 'set-id':
                return {
                    ...state,
                    exerciseId: action.payload,
                }
            case 'set-set':
                return {
                    ...state,
                    sets: action.payload,
                }
            case 'set-rep':
                return {
                    ...state,
                    reps: action.payload,
                }
        }
    }

    const [data, actionDispatch] = useReducer(actionReducer, {
        id: 0,
        exerciseId: 0,
        sets: 0,
        reps: 0,
        available: false,
    })

    const changeValue = (type, val) => {
        switch (type) {
            case 'id':
                actionDispatch({ type: 'set-id', payload: parseInt(val) })
                break
            case 'set':
                actionDispatch({ type: 'set-set', payload: parseInt(val) })
                break
            case 'rep':
                actionDispatch({ type: 'set-rep', payload: parseInt(val) })
                break
        }
    }

    useEffect(() => {
        onChange(data)
    }, [data])

    useEffect(() => {
        const { id, exerciseId, sets, reps } = action
        actionDispatch({
            type: 'set-all',
            payload: { id, exerciseId, sets, reps },
        })
    }, [action])

    return data.available ? (
        <div>
            <ExerciseSelect selected={data.exerciseId} onChange={(e) => changeValue('id', e)} />
            <label htmlFor='reps'>Sets</label>
            <input
                name='reps'
                type='text'
                value={data.sets}
                onChange={(e) => changeValue('set', e.target.value)}
            />
            <label htmlFor='reps'>Reps</label>
            <input
                name='reps'
                type='text'
                value={data.reps}
                onChange={(e) => changeValue('rep', e.target.value)}
            />
        </div>
    ) : null
}

const Actions = ({ actions }) => {
    const onChange = (data) => {
        console.log('data: ', data)
    }

    const addAction = () => {

    }

    return (
        <>
            {actions && actions.length > 0 ? (
                actions.map((a) => (
                    <Action key={a.id} action={a} onChange={onChange} />
                ))
            ) : null}
            <button type='button' onClick={addAction}>Add Exercise</button>
        </>
    )
}

const DateTime = ({ date }) => {
    const d = new Date(date)
    const str = `${d.getMonth() + 1}-${d.getDate()}-${d.getFullYear()}`

    return <div>{str}</div>
}

const Workouts = ({ }) => {
    const [workouts, setWorkouts] = useState([])

    useEffect(() => {
        GetAll('workouts')
            .then((res) => {
                setWorkouts(res)
            })
            .catch((err) => {
                console.log('workouts err: ', err)
            })
    }, [])

    const create = () => {
        GetAll('workouts/create')
            .then((res) => {
                setWorkouts(res)
            })
            .catch((err) => {
                console.log('workouts err: ', err)
            })
    }

    return workouts.length > 0 ? (
        <>
            {workouts.map((w) => (
                <div key={w.id}>
                    <DateTime date={w.date} />
                    <Actions actions={w.workoutActions.actions} />
                </div>
            ))}
            <button type='button' onClick={create}>Add Workout</button>
        </>
    ) : null
}

export default Workouts
