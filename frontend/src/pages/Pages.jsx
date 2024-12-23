import Exercises from './Exercises'
import Workouts from './Workouts'

const Pages = ({ page }) => {
    let Page = null

    switch (page) {
        case 'exercises':
            Page = Exercises
            break
        case 'workouts':
            Page = Workouts
            break
    }

    return <Page />
}

export default Pages
