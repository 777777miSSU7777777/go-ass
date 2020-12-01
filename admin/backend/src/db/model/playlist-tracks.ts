import { Model, snakeCaseMappers } from 'objection';

export class PlaylistTracks extends Model {
    playlistId: string;
    trackId: string;
    
    static get tableName() {
        return 'playlist_tracks';
    }

    static get idColumn() {
        return ['playlist_id', 'track_id'];
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default PlaylistTracks;