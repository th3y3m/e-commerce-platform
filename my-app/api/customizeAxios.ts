import axios from "axios";

// Create the axios instance
const instance = axios.create({
    baseURL: 'http://localhost:8080/',
});

// Request interceptor to dynamically set the Authorization header
instance.interceptors.request.use(
    function (config) {
        const token = localStorage.getItem('token'); // Or wherever you store the token
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }
        return config;
    },
    function (error) {
        return Promise.reject(error);
    }
);

// Response interceptor as usual
instance.interceptors.response.use(
    function (response) {
        return response.data;
    },
    function (error) {
        let res = {
            data: {},
            status: 0,
            headers: {},
        };
        if (error.response) {
            res.data = error.response.data;
            res.status = error.response.status;
            res.headers = error.response.headers;
        } else if (error.request) {
            console.log(error.request);
        } else {
            console.log('Error', error.message);
        }
        return Promise.reject(res);
    }
);

export default instance;

// import axios from "axios";

// axios.defaults.withCredentials = true;

// const instance = axios.create({
//     baseURL: 'http://localhost:8080/',
// });

// instance.interceptors.response.use(
//     function (response) {
//         // Ensure that response.data is always returned, and if not, return an empty object or null
//         return response.data;
//     },
//     function (error) {
//         let res = {
//             data: {},
//             status: 0,
//             headers: {},
//         };
//         if (error.response) {
//             res.data = error.response.data;
//             res.status = error.response.status;
//             res.headers = error.response.headers;
//         } else if (error.request) {
//             console.log(error.request);
//         } else {
//             console.log('Error', error.message);
//         }
//         return Promise.reject(res); // Reject the promise with the error response
//     }
// );

// export default instance;