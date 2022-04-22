<template>
  <user-detail :is-edit="true" />
</template>

<script>
import UserDetail from '../components/RepasswdDetail'
export default {
  name: 'EditForm',
  components: { UserDetail }
}
</script>

