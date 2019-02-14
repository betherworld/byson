import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-ranger-header',
  templateUrl: './ranger-header.component.html',
  styleUrls: ['./ranger-header.component.scss']
})
export class RangerHeaderComponent implements OnInit {
  @Input()
  siteNameBlack: string;

  @Input()
  siteNameGreen: string;


  constructor() { }

  ngOnInit() {
  }

}
