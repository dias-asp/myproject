<script>
  import logo from './assets/images/logo-universal.png'
  import {AddTodo, GetTodos, UpdateTodoStatus, DeleteTodo} from '../wailsjs/go/main/App.js'

  let resultText = "Please enter your name below1 👇"
  let name

  let input
  let active = []
  let done = []

  function add() {
    AddTodo(input)
    input = ''
    reload();
  }

  async function toggle(id, done) {
    console.log('Toggle: ', id, ' ', done, '');
    await UpdateTodoStatus(id, done);
    await reload();
  }

  async function reload() {

    try {
      const todos = await GetTodos();

      if (!Array.isArray(todos)) {
        throw new Error('Вместо списка задач получены некорректные данные');
      }

      // Фильтруем задачи на активные и завершённые
      active = todos.filter(todo => !todo.done);
      done = todos.filter(todo => todo.done);


      console.log(`Активные задачи:`, active);
      console.log(`Выполненные задачи:`, done);

    } catch (err) {
      resultText = err.message || 'Не удалось загрузить задачи';
      console.error(resultText);
    }
  }

  async function del(id){
    await DeleteTodo(id);
    await reload();
  }


  reload();
</script>

<main>
<!--  <img alt="Wails logo" id="logo" src="{logo}">-->
  <div class="main">
    <div class="input-box" id="input">
      <input autocomplete="off" bind:value={input} class="input" id="name" type="text" />
      <button class="btn" on:click={add}>Add</button>
    </div>
    <div class="lists">
      <ul id="active">
        Активные
        {#each active as todo}
          <li>
            <div class="item-active">
              <input type="checkbox" class="checkbox" checked = "{false}" on:change={() => toggle(todo.id, 1)}/>
              <span class="text">{todo.content}</span>
              <button class="btn-del" on:click={() => del(todo.id)}>Delete</button>
            </div>
          </li>
        {/each}
      </ul>
      <ul id="done">
        Выполненные
        {#each done as todo}
          <li>
            <div class="item-done">
              <input type="checkbox" class="checkbox" checked="{true}"  on:change={() => toggle(todo.id, 0)}/>
              <span class="text">{todo.content}</span>
              <button class="btn-del" on:click={() => del(todo.id)}>Delete</button>
            </div>
          </li>
        {/each}
      </ul>
    </div>
  </div>



  <button class="btn" on:click={reload}>Reload</button>

</main>

<style>

  .lists {
    display: flex;
  }

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }
  .btn-del {

    width: 30px;
    height: 20px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 3px;
    cursor: pointer;
    font-size: 9px;
  }


  .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    margin-top: 10px;
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .item-done {
    text-decoration: line-through;
    color: gray;
  }

</style>
