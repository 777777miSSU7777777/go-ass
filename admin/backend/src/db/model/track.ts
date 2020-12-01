import { Model, snakeCaseMappers } from 'objection';

export class Track extends Model {
    trackId: string;
    trackTitle: string;
    artistId: string;
    genreId: string;
    uploadedById: string;
    
    static get tableName() {
        return 'tracks';
    }

    static get idColumn() {
        return 'track_id';
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default Track;