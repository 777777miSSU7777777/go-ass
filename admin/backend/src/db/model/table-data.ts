import Artist from "./artist";
import Genre from "./genre";
import GenreTracks from "./genre-tracks";
import Playlist from "./playlist";
import PlaylistTracks from "./playlist-tracks";
import Track from "./track";
import User from "./user";
import UserPlaylists from "./user-playlists";
import UserTokens from "./user-tokens";
import UserTracks from "./user-tracks";

export type TableData =
    | Artist
    | GenreTracks
    | Genre
    | PlaylistTracks
    | Playlist
    | Track
    | UserPlaylists
    | UserTokens
    | UserTracks
    | User;