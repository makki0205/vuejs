import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://localhost:4000',
});

const userService = {
    create(email, password) {
        return instance.post('/api/signup', {
            email: email,
            password: password
        })
    }
}

export default userService