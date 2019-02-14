import {AfterViewInit, Component, OnInit} from '@angular/core';
import { ActivatedRoute } from "@angular/router";


const google: any = (<any> window).google;


@Component({
  selector: 'app-bison-dashboard',
  templateUrl: './bison-dashboard.component.html',
  styleUrls: ['./bison-dashboard.component.scss']
})
export class BisonDashboardComponent implements OnInit {

  BISONS = {
    '1': {
      id: '1',
      name: 'Mihaela',
      image: 'token_bison1.png',
      icon: 'token_bison1_icon.png',
      age: 3,
      area: 'Carpathians',
    },
    '2': {
      id: '2',
      name: 'Andrei',
      image: 'token_bison2.png',
      icon: 'token_bison2_icon.png',
      age: 5,
      area: 'Carpathians',
    },
    '3': {
      id: '3',
      name: 'Mihai',
      image: 'token_bison3.png',
      icon: 'token_bison3_icon.png',
      age: 7,
      area: 'Carpathians',
    },
  };

  bison;
  lat: number = 47.108090;
  lng: number = 25.530117;


  constructor(private route: ActivatedRoute) {}


  async ngOnInit() {
    this.route.params.subscribe(params => {
      this.bison = this.BISONS[params.bison_id];
    });
  }
}
