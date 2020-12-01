import { Model, snakeCaseMappers } from 'objection';

export class UserTracks extends Model {
    userId: string;
    trackId: string;
    
    static get tableName() {
        return 'user_tracks';
    }

    static get idColumn() {
        return ['user_id', 'track_id'];
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default UserTracks;