import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-bison-dashboard',
  templateUrl: './bison-dashboard.component.html',
  styleUrls: ['./bison-dashboard.component.scss']
})
export class BisonDashboardComponent implements OnInit {
  name;
  balance;


  async ngOnInit() {
    const bison = await (await fetch("http://localhost:10050/bison")).json();
    console.log(bison);

    this.name = bison.name;
  }

}
