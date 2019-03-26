export default class NetworkHandler {
  public static async Init() {
    const url = "http://" + self.location.host + "/init";
    console.log(url)
    return fetch(url, {
      method: 'POST',
    }).then(this.checkStatus)
      .then(response => {
        console.log("Success")
        return response;
      })
      .catch(e => {
        return e;
      });
  }

  public static checkStatus(response: Response) {
    if (response.ok) {
      return response;
    } 
      console.log('There has been an error');
      throw response;
    
  }
}