import { Component, Input, OnInit } from '@angular/core';
import { ListItem } from 'src/app/services/interfaces/list';

@Component({
  selector: 'app-list-item',
  templateUrl: './list-item.component.html',
  styleUrls: ['./list-item.component.scss'],
})
export class ListItemComponent implements OnInit {
  @Input() data: ListItem | undefined;

  starsClass: string[] = [];
  constructor() {}

  ngOnInit(): void {
    if (this.data) {
      for (let i = 0; i < 5; i++) {
        if (i < this.data.usersRating) this.starsClass.push('star-on');
        else this.starsClass.push('star-off');
      }
    }
  }
}
