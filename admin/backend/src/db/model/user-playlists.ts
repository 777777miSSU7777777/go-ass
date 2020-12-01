import { Model, snakeCaseMappers } from 'objection';

export class UserPlaylists extends Model {
    userId: string;
    playlistId: string;
    
    static get tableName() {
        return 'user_playlists';
    }

    static get idColumn() {
        return ['user_id', 'playlist_id'];
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default UserPlaylists;