import { Model, snakeCaseMappers } from 'objection';

export class Playlist extends Model {
    playlistId: string;
    playlistTitle: string;
    createdById: string;

    static get tableName() {
        return 'playlists';
    }

    static get idColumn() {
        return 'playlist_id';
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default Playlist;