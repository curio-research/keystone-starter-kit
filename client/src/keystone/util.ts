import axios, { AxiosInstance } from 'axios';

// Keystone API
export class KeystoneAPIBase {
  private keystoneServerURL: string;
  private keystoneWebsocketURL: string;
  private httpAPI: AxiosInstance;

  constructor(keystoneServerURL: string, keystoneWebsocketURL: string) {
    this.keystoneServerURL = keystoneServerURL;
    this.keystoneWebsocketURL = keystoneWebsocketURL;

    const keystoneAPI = axios.create({
      baseURL: keystoneServerURL, // Replace with your API's base URL
    });

    this.httpAPI = keystoneAPI;
  }

  // Get Keystone server URL
  public getServerURL = (): string => {
    return this.keystoneServerURL;
  };

  // Get Keystone websocket URL
  public getWebsocketURL = (): string => {
    return this.keystoneWebsocketURL;
  };

  public getAPI = (): AxiosInstance => {
    return this.httpAPI;
  };
}
