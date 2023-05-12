import axiosConfig from '../config/axiosConfig';

export const getAllImages = () => {
    var data = axiosConfig.get('/image')
    return data;

}

export const getImageById = (id) => {
    var data = axiosConfig.get(`/image/${id}`)

    return data;

}
