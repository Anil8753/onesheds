import {
  Component,
  Input,
  OnChanges,
  OnInit,
  SimpleChanges,
} from '@angular/core';

@Component({
  selector: 'app-text-list',
  templateUrl: './text-list.component.html',
  styleUrls: ['./text-list.component.scss'],
})
export class TextListComponent implements OnInit, OnChanges {
  //
  @Input() label: string = '';
  @Input() items: string[] = [];
  @Input() max: number = 5;

  constructor() {}

  ngOnChanges(changes: SimpleChanges): void {
    const change = changes['items'];

    if (change.currentValue != change.previousValue) {
    }
  }

  ngOnInit(): void {}
}
