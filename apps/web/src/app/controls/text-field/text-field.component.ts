import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-text-field',
  templateUrl: './text-field.component.html',
  styleUrls: ['./text-field.component.scss'],
})
export class TextFieldComponent implements OnInit {
  //
  @Input() label: string = '';
  @Input() text: string = '';

  width = 200;

  constructor() {}

  ngOnInit(): void {}
}
