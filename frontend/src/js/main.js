"use strict";

var audioTrackList;
var currentAuthor;
var currentTitle;
var player;
var hls;
var searchField;
var searchButton;
var openAddAudioModalButton;
var addAudioModal;
var closeAudioModalButton;
var addAudioButton;

window.onload = () => {
    audioTrackList = document.getElementById("audio-list-container");
    currentAuthor = document.getElementById("current-author");
    currentTitle = document.getElementById("current-title");
    player = document.getElementById("player");
    hls = new Hls();
    searchField = document.getElementById("search-field");
    searchButton = document.getElementById("search-button");
    openAddAudioModalButton = document.getElementById("open-audio-form-modal-button");
    addAudioModal = document.getElementById("add-audio-modal");
    closeAudioModalButton = document.getElementById("close-audio-modal");
    addAudioButton = document.getElementById("add-audio-button");

    player.onpause = () => {
        let controlButton = document.getElementById(getCurrentTrackID()).firstElementChild;
        controlButton.className = "button play";
        controlButton.removeEventListener("click", pauseAudio);
        controlButton.addEventListener("click", resumeAudio);
    };

    player.onplaying = () => {
        let controlButton = document.getElementById(getCurrentTrackID()).firstElementChild;
        controlButton.className = "button pause";
        controlButton.removeEventListener("click", resumeAudio);
        controlButton.addEventListener("click", pauseAudio);
    };
    
    searchButton.addEventListener("click", searchAudio);

    searchField.addEventListener("keyup", e => {
        if (e.keyCode == 13) {
            e.preventDefault();
            searchAudio();
        }
    });

    player.onended = () => {
        nextTrack(getCurrentTrackID());
    };

    document.addEventListener("keyup", e => {
        if (e.keyCode == 36){
            e.preventDefault();
            prevTrack(getCurrentTrackID());
        };

        if (e.keyCode == 33){
            e.preventDefault();
            nextTrack(getCurrentTrackID());
        };

        if (e.keyCode == 38){
            e.preventDefault();
            if (player.paused){
                player.play();
            } else {
                player.pause();
            }
        };

        if (e.keyCode == 107){
            e.preventDefault();
            player.volume += 0.05;
        };
        if (e.keyCode == 13){
            e.preventDefault();
            player.volume -= 0.05;
        };
    });

    openAddAudioModalButton.onclick = () => {
        addAudioModal.style.display = "block";
    };

    closeAudioModalButton.onclick = () => {
        addAudioModal.style.display = "none";
    };

    window.onclick = e => {
        if (e.target == addAudioModal) {
            addAudioModal.style.display = "none";
        }
    }

    addAudioButton.onclick = e => {
        let author = document.getElementById("audio-author-field").value;
        document.getElementById("audio-author-field").value = "";
        let title = document.getElementById("audio-title-field").value;
        document.getElementById("audio-title-field").value = "";
        let audioFile = document.getElementById("audio-file-field").files[0];
        document.getElementById("audio-file-field").value = "";

        let formData = new FormData();
        formData.append("author", author);
        formData.append("title", title);
        formData.append("audiofile", audioFile);

        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/api/audio");
        xhr.send(formData);
        xhr.onload = xhr.onerror = () => {
            fetch("/api/audio", { method: "GET"})
            .then(resp => resp.json())
            .then(data => renderAudioList(data["audio"]));
        }
        addAudioModal.style.display = "none";
    }

    player.onkeydown = e => {
        if (e.keyCode == 37){
            rewind();
        }
        if (e.keyCode == 39){
            flashForward();
        }
    }

    fetch("/api/audio", { method: "GET"})
        .then(resp => resp.json())
        .then(data => renderAudioList(data["audio"]));
};

const createAudioTrackElement = (id, author, title) => {
    let audioTrackElement = document.createElement("div");
    audioTrackElement.className = "audio-track";
    audioTrackElement.id = id;
    
    let controlButton = document.createElement("div");
    controlButton.className = "button play";
    controlButton.addEventListener("click", playAudio);
    
    let authorElement = document.createElement("p");
    authorElement.className = "author";
    authorElement.innerText = author;

    let delimiterElement = document.createElement("p");
    delimiterElement.innerText = " - ";

    let titleElement = document.createElement("p");
    titleElement.className = "title";
    titleElement.innerText = title;

    let rightControls = document.createElement("div");
    rightControls.className = "right-controls";

    let downloadButton = document.createElement("div");
    downloadButton.className = "button download";
    downloadButton.addEventListener("click", downloadAudio);

    let deleteButton = document.createElement("div");
    deleteButton.className = "button delete";
    deleteButton.addEventListener("click", deleteAudio);

    audioTrackElement.appendChild(controlButton);
    audioTrackElement.appendChild(authorElement);
    audioTrackElement.appendChild(delimiterElement);
    audioTrackElement.appendChild(titleElement);
    rightControls.appendChild(downloadButton);
    rightControls.appendChild(deleteButton);
    audioTrackElement.appendChild(rightControls);

    return audioTrackElement;
};

const setCurrentTrackID = id => {
    localStorage.setItem("currentTrackID", id);
};

const getCurrentTrackID = () => {
    return localStorage.getItem("currentTrackID")
};

const loadAudio = (id, author, title) => {
    changeCurrentTrack(author, title);
    setCurrentTrackID(id);
    hls.loadSource("/media/" + id + "/stream/");
    hls.attachMedia(player);
};

const prevTrack = id => {
    let currentTrack = document.getElementById(id);
    let prevTrack = currentTrack.previousSibling;
    if (prevTrack == null) {
        prevTrack = audioTrackList.lastElementChild;
    }
    let prevID = prevTrack.id;
    let author = prevTrack.getElementsByClassName("author")[0].innerText;
    let title = prevTrack.getElementsByClassName("title")[0].innerText;

    setCurrentTrackID(prevID);
    renderTracksControls();
    loadAudio(prevID, author, title);
    player.play();
};

const nextTrack = id => {
    let currentTrack = document.getElementById(id);
    let nextTrack = currentTrack.nextSibling;
    if (nextTrack == null) {
        nextTrack = audioTrackList.firstElementChild;
    }
    let nextID = nextTrack.id;
    let author = nextTrack.getElementsByClassName("author")[0].innerText;
    let title = nextTrack.getElementsByClassName("title")[0].innerText;

    setCurrentTrackID(nextID);
    renderTracksControls();
    loadAudio(nextID, author, title);
    player.play();
};

const playAudio = e => {
    let id = e.target.parentNode.id;
    let author = e.target.parentNode.getElementsByClassName("author")[0].innerText;
    let title = e.target.parentNode.getElementsByClassName("title")[0].innerText;
    loadAudio(id, author, title)
    player.play();
    
    let tracksList = audioTrackList.children;
    for (let i = 0; i < tracksList.length; i++){
        tracksList[i].firstElementChild.className = "button play";
        tracksList[i].firstElementChild.removeEventListener("click", resumeAudio);
        tracksList[i].firstElementChild.removeEventListener("click", pauseAudio);
        tracksList[i].firstElementChild.addEventListener("click", playAudio);
    };

    e.target.removeEventListener("click", playAudio);
    e.target.addEventListener("click", pauseAudio);
};

const pauseAudio = () => {
    player.pause();
};

const resumeAudio = () => {
    player.play();
};

const changeCurrentTrack = (author, title) => {
    currentAuthor.innerText = author;
    currentTitle.innerText = title;
    document.title = author + " - " + title;
};

const renderAudioList = (tracks) => {
    var audioTracksElements = tracks.map(v => createAudioTrackElement(v.id, v.author, v.title));

    while (audioTrackList.firstChild) {
        audioTrackList.removeChild(audioTrackList.firstChild);
    };
 
    audioTracksElements.forEach(v => audioTrackList.appendChild(v));
};

const renderTracksControls = () => {
    let tracksList = audioTrackList.children;
    for (let i = 0; i < tracksList.length; i++){
        tracksList[i].firstElementChild.className = "button play";
    }
    document.getElementById(getCurrentTrackID()).firstElementChild.className = "button pause";
};

const searchAudio = () => {
    let searchKey = searchField.value;
    fetch("/api/audio?key=" + searchKey, {method: "GET"})
        .then(resp => resp.json())
        .then(data => {
            renderTracksControls();
            renderAudioList(data["audio"])
        });
};

const rewind = () => {
    player.currentTime -= 10;
}

const flashForward = () => {
    player.currentTime += 10;
}

const deleteAudio = e => {
    let track = e.target.parentNode.parentNode;
    let id = track.id;
    let xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/api/audio/" + id);
    xhr.send();

    xhr.onload = () => {
        if (xhr.status == 200) {
            track.remove();
        }
    };
}

const downloadAudio = e => {
    let track = e.target.parentNode.parentNode;
    let id = track.id;
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/media/" + id + "/download");
    xhr.send();

    xhr.onload = () => {
        if (xhr.status == 200){
            var a = document.createElement("a");
            let author = track.getElementsByClassName("author")[0].innerText;
            let title = track.getElementsByClassName("title")[0].innerText;
            a.href = xhr.responseURL;
            a.setAttribute("download", author + " - " + title);
            a.click();
        }
    }
}