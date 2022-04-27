import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { DataService } from 'src/app/services/data.service';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss'],
})
export class DetailsComponent implements OnInit {
  //
  whId: any;
  data: any;

  constructor(
    public dialogRef: MatDialogRef<DetailsComponent>,
    @Inject(MAT_DIALOG_DATA) public input: DialogData,
    private _dataservice: DataService
  ) {
    this.whId = input.whId;
  }

  ngOnInit(): void {
    const promise = this._dataservice.detailService.get(this.whId);
    promise.then((r) => {
      this.data = r;
    });
  }

  close(): void {
    this.dialogRef.close();
  }
}

export interface DialogData {
  whId: string;
}
function e(e: any) {
  throw new Error('Function not implemented.');
}
