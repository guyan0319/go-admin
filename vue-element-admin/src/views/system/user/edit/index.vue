<template>
  <user-detail :is-edit="true" />
</template>

<script>
import UserDetail from '../components/EditDetail'
export default {
  name: 'EditForm',
  components: { UserDetail }
}
</script>

