import { DataTables } from "../enums/data-tables"

export const getDataRouteForTable = (table: DataTables): string => {
    switch(+table) {
        case DataTables.artist:
            return '/artist';
        case DataTables.genreTracks:
            return '/genre-tracks';
        case DataTables.genre:
            return '/genre';
        case DataTables.playlistTracks:
            return '/playlist-tracks';
        case DataTables.playlist:
            return '/playlist';
        case DataTables.track:
            return '/track';
        case DataTables.userPlaylists:
            return '/user-playlists';
        case DataTables.userTokens:
            return '/user-tokens';
        case DataTables.userTracks:
            return '/user-tracks';
        case DataTables.user:
            return '/user';
        default:
            return '';
    }
}

export const getTableDataFields = (table: DataTables): { db: string[], obj: string[], pk: string[] } => {
    switch(+table) {
        case DataTables.artist:
            return {
                db: ['artist_id', 'artist_name'],
                obj: ['artistId', 'artistName'],
                pk: ['artistId'],
            };
        case DataTables.genreTracks:
            return {
                db: ['genre_id', 'track_id'],
                obj: ['genreId', 'trackId'],
                pk: ['genreId', 'trackId'],
            };
        case DataTables.genre:
            return {
                db: ['genre_id', 'genre_title'],
                obj: ['genreId', 'genreTitle'],
                pk: ['genreId'],
            };
        case DataTables.playlistTracks:
            return {
                db: ['playlist_id', 'track_id'],
                obj: ['playlistId', 'trackId'],
                pk: ['playlistId', 'trackId']
            };
        case DataTables.playlist:
            return {
                db: ['playlist_id', 'playlist_title', 'created_by_id'],
                obj: ['playlistId', 'playlistTitle', 'createById'],
                pk: ['playlistId'],
            };
        case DataTables.track:
            return {
                db: ['track_id', 'track_title', 'artist_id', 'genre_id', 'uploaded_by_id'],
                obj: ['trackId', 'trackTitle', 'artistId', 'genreId', 'uploadedById'],
                pk: ['trackId'],
            };
        case DataTables.userPlaylists:
            return {
                db: ['user_id', 'playlist_id'],
                obj: ['userId', 'playlistId'],
                pk: ['userId', 'playlistId'],
            };
        case DataTables.userTokens:
            return {
                db: ['user_id', 'token'],
                obj: ['userId', 'token'],
                pk: ['userId', 'token']
            };
        case DataTables.userTracks:
            return {
                db: ['user_id', 'track_id'],
                obj: ['userId', 'trackId'],
                pk:['userId', 'trackId'],
            };
        case DataTables.user:
            return {
                db: ['user_id', 'role', 'email', 'username', 'password'],
                obj: ['userId', 'role', 'email', 'username', 'password'],
                pk: ['userId']
            };

        default:
            return {
                db: [],
                obj: [],
                pk: [],
            };
    }
}

export const getTableName = (table: DataTables): string => {
    switch(+table) {
        case DataTables.artist:
            return 'artists';
        case DataTables.genreTracks:
            return 'genre_tracks';
        case DataTables.genre:
            return 'genres';
        case DataTables.playlistTracks:
            return 'playlist_tracks';
        case DataTables.playlist:
            return 'playlists';
        case DataTables.track:
            return 'tracks';
        case DataTables.userPlaylists:
            return 'user_playlists';
        case DataTables.userTokens:
            return 'user_tokens';
        case DataTables.userTracks:
            return 'user_tracks';
        case DataTables.user:
            return 'users';
        default:
            return '';
    }
}