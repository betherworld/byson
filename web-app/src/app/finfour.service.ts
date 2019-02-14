import { Injectable } from '@angular/core';


const HOST = 'http://www.finfour.net';


@Injectable({
  providedIn: 'root'
})
export class FinfourService {

  constructor() { }


  async login(email, password) {
    const res = await fetch(HOST + '/wapi/login', {
      credentials: 'include',
      method: 'POST',
      body: JSON.stringify({
        email, password
      }),
      mode: 'no-cors',
    });
    console.log(res);
  }
}
