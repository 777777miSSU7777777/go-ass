import { Body, Controller, HttpStatus, Param, Post, Query, Res } from '@nestjs/common';
import { TableDataService } from '@service';
import { TableData } from '@model';
import { Response } from 'express';
import { DataActions } from '@app/enums';

@Controller('/data/:dataRoute')
export class TableDataController {
  constructor(private readonly tableDataService: TableDataService) {}

  @Post()
  async tableAction(@Param() params, @Query() query, @Body() body: TableData[], @Res() res: Response) {
    try {
      const dataRoute: string = params.dataRoute;
      const dataAction: string = query.action;
      let tableData: TableData[];

      switch(dataAction) {
        case DataActions.get:
          tableData = await this.tableDataService.getTableData(dataRoute);
          break;
        case DataActions.create:
          tableData = await this.tableDataService.newTableData(dataRoute, body);
          break;
        case DataActions.update:
          tableData = await this.tableDataService.updateTableData(dataRoute, body);
          break;
        case DataActions.delete:
          const deletedRowsCount: number[] = await this.tableDataService.deleteTableData(dataRoute, body);
          const affectedRowsCount: number = deletedRowsCount.length > 0 ? deletedRowsCount.reduce((acc: number = 0, cur: number) => acc + cur) : 0;
          res.status(HttpStatus.OK).json({
            'statusCode': HttpStatus.OK,
            'ok': true,
            'data': `Affected rows count: ${affectedRowsCount}`,
            'error': null,
          });
          return;
        default:
          throw new Error(`Unknown Data Action: ${dataAction}`);
      }

      res.status(HttpStatus.OK).json({
        'statusCode': HttpStatus.OK,
        'ok': true,
        'data': tableData,
        'error': null,
      });
    } catch (e) {
      res.status(HttpStatus.BAD_REQUEST).json({
        'statusCode': HttpStatus.BAD_REQUEST,
        'ok': false,
        'data': null,
        'error': e,
      });

      console.error(`Table Data Action Error: ${e}`);
    }
  }
}

export default TableDataController;