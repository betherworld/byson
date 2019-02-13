import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-donor-header',
  templateUrl: './donor-header.component.html',
  styleUrls: ['./donor-header.component.scss']
})
export class DonorHeaderComponent implements OnInit {
  isCollapsed = true;


  constructor() { }


  ngOnInit() {
  }

}
