<template>
  <user-detail :is-edit="false" />
</template>

<script>
import UserDetail from '../components/CreateDetail'

export default {
  name: 'CreateArticle',
  components: { UserDetail }
}
</script>

