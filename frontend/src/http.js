import axios from 'axios'

const api_url = `http://localhost:8080/`

///////////////////////////////////////////////////////////
// GET based requests                                    //
///////////////////////////////////////////////////////////

export const GetAll = async (url) => {
    return new Promise((resolve, reject) => {
        axios
            .get(`${api_url}${url}`)
            .then((res) => {
                return resolve(res.data)
            })
            .catch((err) => {
                return reject(err)
            })
    })
}

export const Get = async (url, id) => {
    return new Promise((resolve, reject) => {
        axios
            .get(`${api_url}${url}/${id}`)
            .then((res) => {
                return resolve(res.data)
            })
            .catch((err) => {
                return reject(err)
            })
    })
}

///////////////////////////////////////////////////////////
// POST, PATCH, PUT, DELETE based requests               //
///////////////////////////////////////////////////////////

//create new object
export const Create = async (url, object) => {
    return new Promise((resolve, reject) => {
        axios
            .post(`${api_url}${url}`, object)
            .then((res) => {
                return resolve(res.data.url)
            })
            .catch((err) => {
                return reject(err)
            })
    })
}

//update object
export const Update = async (url, object, id = null) => {
    return new Promise((resolve, reject) => {
        axios
            .patch(`${api_url}${url}/${id}`, object)
            .then((res) => {
                return resolve(res.data)
            })
            .catch((err) => {
                return reject(err)
            })
    })
}

//delete object
export const Delete = async (url, id) => {
    return new Promise((resolve, reject) => {
        axios
            .delete(`${api_url}${url}/${id}`)
            .then((res) => {
                return resolve(res.data)
            })
            .catch((err) => {
                return reject(err)
            })
    })
}


