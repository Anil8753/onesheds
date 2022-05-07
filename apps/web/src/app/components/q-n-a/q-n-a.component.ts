import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Observable } from 'rxjs';
import { map, startWith } from 'rxjs/operators';
import { of } from 'rxjs';

import { DataService } from 'src/app/services/data.service';
import { QnA } from 'src/app/services/interfaces/detail';

@Component({
  selector: 'app-q-n-a',
  templateUrl: './q-n-a.component.html',
  styleUrls: ['./q-n-a.component.scss'],
})
export class QNAComponent implements OnInit {
  formControl = new FormControl();
  entries!: Observable<QnA[]>;
  progress = false;

  constructor(private _dataservice: DataService) {}

  ngOnInit() {
    this.fetch('');

    this.formControl.valueChanges.subscribe(
      (value) => {
        this.fetch(value);
      },
      (err) => {}
    );
  }

  async fetch(word: string) {
    this.progress = true;
    this._dataservice.detailService
      .getQnA('', word)
      .then((res) => {
        this.progress = false;
        this.entries = of(res);
      })
      .catch((e) => {
        this.progress = false;
        console.error(e);
      });
  }
}
