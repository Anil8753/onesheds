import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-search-results-header',
  templateUrl: './search-results-header.component.html',
  styleUrls: ['./search-results-header.component.scss'],
})
export class SearchResultsHeaderComponent implements OnInit {
  //
  @Input() resultsCount = 0;
  @Input() distance = 0;
  @Input() address = '';

  constructor() {}

  ngOnInit(): void {}
}
