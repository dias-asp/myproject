export namespace main {
	
	export class Todo {
	    id: number;
	    content: string;
	    done: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.content = source["content"];
	        this.done = source["done"];
	    }
	}

}

