export namespace config {
	
	export class JavaSettings {
	    preferredPath?: string;
	    maxMemoryMb: number;
	    minMemoryMb: number;
	    extraArgs: string[];
	
	    static createFrom(source: any = {}) {
	        return new JavaSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.preferredPath = source["preferredPath"];
	        this.maxMemoryMb = source["maxMemoryMb"];
	        this.minMemoryMb = source["minMemoryMb"];
	        this.extraArgs = source["extraArgs"];
	    }
	}
	export class Settings {
	    schemaVersion: number;
	    pingIntervalMinutes: number;
	    java: JavaSettings;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.schemaVersion = source["schemaVersion"];
	        this.pingIntervalMinutes = source["pingIntervalMinutes"];
	        this.java = this.convertValues(source["java"], JavaSettings);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace repository {
	
	export class Repository {
	    id: string;
	    name: string;
	    url: string;
	    description: string;
	    enabled: boolean;
	    priority: number;
	
	    static createFrom(source: any = {}) {
	        return new Repository(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.url = source["url"];
	        this.description = source["description"];
	        this.enabled = source["enabled"];
	        this.priority = source["priority"];
	    }
	}

}

