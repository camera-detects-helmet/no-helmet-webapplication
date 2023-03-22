// First we need to import axios.js
import axios from 'axios';


// Next we make an 'instance' of it
const instance = axios.create({
    // .. where we make our configurations
    baseURL: process.env.API_ADDRESS,
});

// instance.defaults.headers.common['Authorization'] = 'AUTH TOKEN FROM INSTANCE';


// instance.interceptors.request.

export default instance;