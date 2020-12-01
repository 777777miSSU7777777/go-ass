import { Model, snakeCaseMappers } from 'objection';

export class Artist extends Model {
    artistId: string;
    artistName: string;
    
    static get tableName() {
        return 'artists';
    }

    static get idColumn() {
        return 'artist_id';
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default Artist;