import { isUndefined } from 'lodash-es/isUndefined';
import { isNull } from 'lodash-es/isNull';
var Config = require('Config');

class ApiClient {
    constructor(baseUrl) {
        this.baseUrl = baseUrl;
    }

    static _instance;   

    static instance() {
        if (!ApiClient._instance){
            ApiClient._instance = new ApiClient(Config.serverUrl);
        }

        return ApiClient._instance;
    }

    getAllAudio() {
       return fetch(`${this.baseUrl}/api/audio`, {method: 'GET'})
                    .then(resp => resp.json())
                    .then(data => data);
    }

    searchAudioByKey(key) {
        return fetch(`${this.baseUrl}/api/audio?key=${key}`, {method: 'GET'})
                    .then(resp => resp.json())
                    .then(data => data['audio']);
    }

    newAudio(form) {
        const formData = new FormData();
        formData.append('author', form.author);
        formData.append('title', form.title);
        formData.append('audiofile', form.file);
        return fetch(`${this.baseUrl}/api/audio`, {method: 'POST', body: formData})
                    .then(resp => resp.json());
    }

    deleteAudioById(id) {
        return fetch(`${this.baseUrl}/api/audio/${id}`, {method: 'DELETE' });
    }
}

export default ApiClient;