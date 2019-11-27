import { isUndefined } from 'lodash-es/isUndefined';
import { isNull } from 'lodash-es/isNull';

class ApiClient {
    constructor(baseUrl){
        this.baseUrl = baseUrl;
    }

    static _instance;   

    static instance(){
        if (!ApiClient._instance){
            ApiClient._instance = new ApiClient("");
        }

        return ApiClient._instance;
    }

    getAllAudio(){
       return fetch(this.baseUrl + "/api/audio", {method: "GET"})
                    .then(resp => resp.json())
                    .then(data => data["audio"]);
    }

    searchAudioByKey(key){
        return fetch(this.baseUrl + "/api/audio?key=" + key, {method: "GET"})
                    .then(resp => resp.json())
                    .then(data => data["audio"]);
    }

    newAudio(formData){
        return fetch(this.baseUrl + "/api/audio", {method: "POST", body: formData})
                    .then(resp => resp.json());
    }

    deleteAudioById(id){
        return fetch(this.baseUrl + "/api/audio/" + id, {method: "DELETE"});
    }

    
}

export default ApiClient;