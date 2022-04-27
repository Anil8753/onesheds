import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ListItem } from 'src/app/services/interfaces/list';

import { DetailsComponent } from '../../details/details.component';

@Component({
  selector: 'app-list-item',
  templateUrl: './list-item.component.html',
  styleUrls: ['./list-item.component.scss'],
})
export class ListItemComponent implements OnInit {
  @Input() data: ListItem | undefined;

  starsClass: string[] = [];
  constructor(private _dialog: MatDialog) {}

  ngOnInit(): void {
    if (this.data) {
      for (let i = 0; i < 5; i++) {
        if (i < this.data.usersRating) this.starsClass.push('star-on');
        else this.starsClass.push('star-off');
      }
    }
  }

  openDetailsDialog(): void {
    const dialogRef = this._dialog.open(DetailsComponent, {
      width: '100vw',
      maxWidth: '100vw',
      height: '100vh',
      maxHeight: '100vh',
      data: { whId: this.data?.id },
    });

    dialogRef.afterClosed().subscribe((result) => {
      console.log('The dialog was closed');
      console.log(result);
    });
  }
}
