import { ArtistService, GenreService, GenreTracksService, PlaylistService, PlaylistTracksService, TableDataService, TrackService, UserPlaylistsService, UserService, UserTokensService, UserTracksService } from '@app/service';
import { Module } from '@nestjs/common';
import { TableDataController } from './table-data.controller';

@Module({
  imports: [],
  controllers: [TableDataController],
  providers: [TableDataService, ArtistService, GenreService, GenreTracksService, PlaylistService, PlaylistTracksService, TrackService, UserService, UserPlaylistsService, UserTokensService, UserTracksService],
})
export class TableDataModule {}
