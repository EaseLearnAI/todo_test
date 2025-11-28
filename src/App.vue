<template>
  <div class="container">
    <h1>时间 Todo 应用</h1>
    
    <form @submit.prevent="createTodo" class="card">
      <div class="field">
        <label>标题</label>
        <input 
          type="text" 
          v-model="newTodo.title" 
          placeholder="例如：完成周报" 
          required 
        />
      </div>
      <div class="field">
        <label>截止时间</label>
        <input 
          type="datetime-local" 
          v-model="newTodo.dueAt" 
        />
      </div>
      <button type="submit">添加</button>
    </form>

    <div class="toolbar">
      <button @click="fetchTodos">刷新</button>
    </div>

    <ul class="list">
      <li 
        v-for="todo in sortedTodos" 
        :key="todo.id" 
        :class="['item', { done: todo.completed }]"
      >
        <div class="left">
          <input 
            type="checkbox" 
            :checked="todo.completed" 
            @change="toggleTodo(todo.id)" 
          />
          <input 
            type="text" 
            v-model="todo.title" 
            class="title-input" 
          />
        </div>
        
        <div class="right">
          <input 
            type="datetime-local" 
            v-model="todo.dueAtLocal" 
            class="due-input" 
          />
          <span class="due-label">{{ formatDate(todo.dueAt) }}</span>
          <button @click="updateTodo(todo)">保存</button>
          <button @click="deleteTodo(todo.id)" class="danger">删除</button>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const todos = ref([])
const newTodo = ref({
  title: '',
  dueAt: ''
})

// API 调用
const api = {
  async list() {
    const res = await fetch('/api/todos')
    const json = await res.json()
    return json.data || []
  },
  async create(payload) {
    const res = await fetch('/api/todos', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })
    return res.json()
  },
  async update(id, payload) {
    const res = await fetch(`/api/todos/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })
    return res.json()
  },
  async toggle(id) {
    const res = await fetch(`/api/todos/${id}/toggle`, { 
      method: 'POST' 
    })
    return res.json()
  },
  async remove(id) {
    const res = await fetch(`/api/todos/${id}`, { 
      method: 'DELETE' 
    })
    return res.json()
  }
}

// 格式化日期
const formatDate = (iso) => {
  if (!iso) return '未设置'
  const d = new Date(iso)
  if (isNaN(d.getTime())) return '时间异常'
  return d.toLocaleString('zh-CN')
}

// 转换 ISO 时间
const isoFromLocalInput = (value) => {
  if (!value) return null
  const d = new Date(value)
  if (isNaN(d.getTime())) return null
  return d.toISOString()
}

// 转换为本地时间格式
const localInputFromIso = (iso) => {
  if (!iso) return ''
  try {
    const d = new Date(iso)
    const pad = (n) => String(n).padStart(2, '0')
    return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}T${pad(d.getHours())}:${pad(d.getMinutes())}`
  } catch {
    return ''
  }
}

// 排序后的 Todos
const sortedTodos = computed(() => {
  return [...todos.value]
    .map(t => ({
      ...t,
      dueAtLocal: localInputFromIso(t.dueAt)
    }))
    .sort((a, b) => {
      const aDue = a.dueAt ? new Date(a.dueAt).getTime() : Infinity
      const bDue = b.dueAt ? new Date(b.dueAt).getTime() : Infinity
      if (a.completed !== b.completed) return a.completed ? 1 : -1
      return aDue - bDue
    })
})

// 获取所有 Todos
const fetchTodos = async () => {
  todos.value = await api.list()
}

// 创建 Todo
const createTodo = async () => {
  const payload = {
    title: newTodo.value.title.trim(),
    dueAt: isoFromLocalInput(newTodo.value.dueAt)
  }
  
  const res = await api.create(payload)
  if (res.error) {
    alert(res.error)
  } else {
    newTodo.value.title = ''
    newTodo.value.dueAt = ''
    await fetchTodos()
  }
}

// 更新 Todo
const updateTodo = async (todo) => {
  const payload = {
    title: todo.title.trim(),
    dueAt: isoFromLocalInput(todo.dueAtLocal)
  }
  
  const res = await api.update(todo.id, payload)
  if (res.error) {
    alert(res.error)
  } else {
    await fetchTodos()
  }
}

// 切换完成状态
const toggleTodo = async (id) => {
  await api.toggle(id)
  await fetchTodos()
}

// 删除 Todo
const deleteTodo = async (id) => {
  await api.remove(id)
  await fetchTodos()
}

// 初始化
onMounted(() => {
  fetchTodos()
})
</script>
