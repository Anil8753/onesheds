import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';

@Component({
  selector: 'app-search-ctrl',
  templateUrl: './search-ctrl.component.html',
  styleUrls: ['./search-ctrl.component.scss'],
})
export class SearchCtrlComponent implements OnInit {
  city = '';
  cities = ['Bengaluru', 'Hyderabad'];

  locality = '';
  localities = [
    'HSR Layout',
    'BTM Layout',
    'Whitefield',
    'Hoskote',
    'Hopefarm',
  ];

  constructor(private _route: Router, private _snackBar: MatSnackBar) {}

  ngOnInit(): void {}

  onSearch() {
    if (this.city === '') {
      this.openSnackBar('Missed', `Please select city and locality`);
      return;
    }

    if (this.locality === '') {
      this.openSnackBar(
        'Missed',
        `Please select locality within city ${this.city}`
      );
      return;
    }

    this._route.navigate(['listing'], {
      queryParams: { locality: this.locality },
    });
  }

  private openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action, {
      duration: 3000,
    });
  }
}
